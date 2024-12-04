import { Container, Flex, Separator, Text } from "@radix-ui/themes";
import { useQuery } from "@tanstack/react-query";
import { Interweave } from "interweave";
import { useAtomValue } from "jotai";
import { useParams } from "react-router-dom";
import { axiosLocalhost } from "../../api/axios/axios";
import { userAtom } from "../../AtomStore/AtomStore";
import ChangePostButton from "./ChangePostButton/ChangePostButton";
import SkeletonPostLoader from "./SkeletonLoader/SkeletonLoader";

type TArticleViewer = {
    htmlToParse?: string;
};

export default function ArticleViewer(props: TArticleViewer) {
    let queryParams = useParams();
    const user = useAtomValue(userAtom);

    const { data, isPending } = useQuery({
        queryKey: [`post_${queryParams["postId"]}`],
        queryFn: async () => {
            const response = await axiosLocalhost.get(
                `posts/${queryParams["postId"]}`
            );

            return response.data;
        },
    });

    if (isPending) return <SkeletonPostLoader />;

    return (
        <>
            <Container size={"3"}>
                <div className="ql-snow ql-editor">
                    <Container size={"2"} className="mt-4">
                        <Flex direction={"column"}>
                            <Text className="mb-2" as="div" size={"9"}>
                                {data.title}
                            </Text>
                            <Flex className="mt-4 mb-2">
                                <div hidden={data.user_id != user?.id}>
                                    <ChangePostButton
                                        postId={queryParams["postId"] || ""}
                                    />
                                </div>
                            </Flex>
                        </Flex>
                        <Separator size={"4"} className="mb-2" />
                        <Interweave content={data.content} />
                    </Container>
                </div>
            </Container>
        </>
    );
}
