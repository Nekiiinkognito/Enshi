import {
    Box,
    Container,
    Flex,
    Separator,
    Text
} from "@radix-ui/themes";
import { useQuery } from "@tanstack/react-query";
import { Interweave } from "interweave";
import { useAtomValue } from "jotai";
import { useParams } from "react-router-dom";
import { axiosLocalhost } from "../../api/axios/axios";
import { userAtom } from "../../AtomStore/AtomStore";
import AddPostToBlogDialog from "../Dialogs/AddPostToBlogDialog/AddPostToBlogDialog";
import ChangePostButton from "./ChangePostButton/ChangePostButton";
import SkeletonPostLoader from "./SkeletonLoader/SkeletonLoader";
import VoteButton, { DOWNVOTE, UPVOTE } from "./VoteButton/VoteButton";
import VoteCounter from "./VoteCounter/VoteCounter";

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
        gcTime: 0,
        refetchOnMount: true,
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
                            <Flex
                                gap={"3"}
                                className="items-center mt-4 mb-2 align-baseline"
                            >
                                <Flex gap={"1"}>
                                    <VoteButton
                                        vote={UPVOTE}
                                        postId={queryParams["postId"] || ""}
                                    />

                                    <VoteCounter
                                        postId={queryParams["postId"] || ""}
                                    />

                                    <VoteButton
                                        vote={DOWNVOTE}
                                        postId={queryParams["postId"] || ""}
                                    />
                                </Flex>

                                <Box hidden={data.user_id != user?.id}>
                                    <ChangePostButton
                                        postId={queryParams["postId"] || ""}
                                    />
                                </Box>

                                <AddPostToBlogDialog />
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
