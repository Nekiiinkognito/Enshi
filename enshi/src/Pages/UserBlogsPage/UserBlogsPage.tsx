import { Box, Container, Flex } from "@radix-ui/themes";
import { useQuery } from "@tanstack/react-query";
import { useAtomValue } from "jotai";
import { userAtom } from "../../AtomStore/AtomStore";
import { axiosLocalhost } from "../../api/axios/axios";
import BlogBox from "./BlogBox/BlogBox";
import SkeletonBoxes from "./SkeletonBoxes/SkeletonBoxes";

export default function UserBlogsPage() {
    const user = useAtomValue(userAtom);

    const isBigNumber = (num: any) => !Number.isSafeInteger(+num);

    const enquoteBigNumber = (jsonString: any, bigNumChecker: any) =>
        jsonString.replaceAll(
            /([:\s\[,]*)(\d+)([\s,\]]*)/g,
            (matchingSubstr: any, prefix: any, bigNum: any, suffix: any) =>
                bigNumChecker(bigNum)
                    ? `${prefix}"${bigNum}"${suffix}`
                    : matchingSubstr
        );

    const parseWithBigInt = (jsonString: any, bigNumChecker: any) =>
        JSON.parse(enquoteBigNumber(jsonString, bigNumChecker), (key, value) =>
            !isNaN(value) && bigNumChecker(value)
                ? BigInt(value).toString()
                : value
        );

    const { data, isPending, isFetching } = useQuery({
        queryKey: ["userBlogs"],
        queryFn: async () => {
            const response = await axiosLocalhost.get("/user/blogs", {
                transformResponse: [(data) => data],
            });

            let temp = parseWithBigInt(response.data, isBigNumber);

            return temp as any[];
        },
    });

    if (isFetching)
        return (
            <Container size={"1"}>
                <SkeletonBoxes />
            </Container>
        );

    return (
        <Box className="size-full">
            <Container size={"1"}>
                <Flex direction={"column"} gap={"2"}>
                    {data
                        ? data?.map((blog: any, b) => {
                              return (
                                  <BlogBox
                                      key={b}
                                      title={blog.title}
                                      blogId={blog.blog_id}
                                  />
                              );
                          })
                        : null}
                </Flex>
            </Container>
        </Box>
    );
}
