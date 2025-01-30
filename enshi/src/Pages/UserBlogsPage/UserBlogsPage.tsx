import * as ScrollArea from "@radix-ui/react-scroll-area";
import { Box, Container, Flex, Separator, Text } from "@radix-ui/themes";
import { useQuery } from "@tanstack/react-query";
import { axiosLocalhost } from "../../api/axios/axios";
import BlogBox from "../../Components/BlogBox/BlogBox";
import BlogCreationDialog from "../../Components/Dialogs/BlogCreationDialog/BlogCreationDialog";
import { JSONWithInt64 } from "../../utils/idnex";
import SkeletonBoxes from "./SkeletonBoxes/SkeletonBoxes";

const TAGS = Array.from({ length: 50 }).map(
    (_, i, a) => `v1.2.0-beta.${a.length - i}`
);

export default function UserBlogsPage() {
    const { data, isPending, isFetching } = useQuery({
        queryKey: ["userBlogs"],
        queryFn: async () => {
            const response = await axiosLocalhost.get("/user/blogs", {
                transformResponse: [(data) => data],
            });

            let temp = JSONWithInt64(response.data);

            return temp as any[];
        },
    });

    if (isPending)
        return (
            <Container size={"1"}>
                <SkeletonBoxes />
            </Container>
        );

    return (
        <Box className="w-full max-h-full overflow-hidden">
            {/* <Container size={"2"} className="w-full h-full max-h-full"> */}
            <Flex
                id="currentTestBox"
                direction={"column"}
                gap={"2"}
                className="max-h-full pb-2 mx-80"
            >
                <Text size={"9"} className="text-center">
                    Your blogs
                </Text>

                <Separator size={"4"} className="my-2" />

                <div className="overflow-hidden">
                    <ScrollArea.Root className="w-full overflow-hidden h-fit max-h-">
                        <ScrollArea.Viewport className="size-full">
                            <Flex direction={"column"} gap={"2"}>
                                {data
                                    ? data?.map((blog: any, b) => {
                                          return (
                                              <>
                                                  <BlogBox
                                                      key={b}
                                                      title={blog.title}
                                                      blogId={blog.blog_id}
                                                      userId={blog.user_id}
                                                  />
                                              </>
                                          );
                                      })
                                    : null}
                            </Flex>
                        </ScrollArea.Viewport>
                        <ScrollArea.Scrollbar
                            orientation="vertical"
                            color="gray"
                            className="rounded-xl data-[state=visible]:w-1"
                        >
                            <ScrollArea.Thumb className="relative flex-1 bg-gray-600 rounded-xl" />
                        </ScrollArea.Scrollbar>
                    </ScrollArea.Root>
                </div>

                <BlogCreationDialog />

              
            </Flex>
            {/* </Container> */}
        </Box>
    );
}
