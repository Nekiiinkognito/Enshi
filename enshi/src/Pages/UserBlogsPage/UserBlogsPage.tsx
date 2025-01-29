import {
    Box,
    Container,
    Flex,
    Separator,
    Text
} from "@radix-ui/themes";
import { useQuery } from "@tanstack/react-query";
import { axiosLocalhost } from "../../api/axios/axios";
import BlogBox from "../../Components/BlogBox/BlogBox";
import BlogCreationDialog from "../../Components/Dialogs/BlogCreationDialog/BlogCreationDialog";
import { JSONWithInt64 } from "../../utils/idnex";
import SkeletonBoxes from "./SkeletonBoxes/SkeletonBoxes";

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
        <Box className="size-full">
            <Container size={"1"}>
                <Flex direction={"column"} gap={"2"}>
                    <Text size={"9"} className="text-center">
                        Your blogs
                    </Text>

                    <Separator size={"4"} className="my-2" />

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

                    <BlogCreationDialog />
                </Flex>
            </Container>
        </Box>
    );
}
