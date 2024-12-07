import * as Dialog from "@radix-ui/react-dialog";
import { Cross2Icon } from "@radix-ui/react-icons";
import {
    Box,
    Button,
    Container,
    Flex,
    Select,
    Separator,
    Text,
} from "@radix-ui/themes";
import { useQuery } from "@tanstack/react-query";
import { Interweave } from "interweave";
import { useAtomValue } from "jotai";
import { useParams } from "react-router-dom";
import { axiosLocalhost } from "../../api/axios/axios";
import { userAtom } from "../../AtomStore/AtomStore";
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

                                <Dialog.Root>
                                    <Dialog.Trigger asChild>
                                        <Button
                                            variant="surface"
                                            className="h-5"
                                        >
                                            Add to blog
                                        </Button>
                                    </Dialog.Trigger>
                                    <Dialog.Portal>
                                        <Dialog.Overlay className="fixed inset-0 bg-blackA6 data-[state=open]:animate-overlayShow" />
                                        <Dialog.Content className="fixed left-1/2 top-1/2 max-h-[85vh] w-[90vw] max-w-[450px] -translate-x-1/2 -translate-y-1/2 rounded-md bg-white p-[25px] shadow-[hsl(206_22%_7%_/_35%)_0px_10px_38px_-10px,_hsl(206_22%_7%_/_20%)_0px_10px_20px_-15px] focus:outline-none data-[state=open]:animate-contentShow">
                                            <Dialog.Title className="m-0 text-[17px] font-medium text-mauve12">
                                                Add this post to blog
                                            </Dialog.Title>
                                            <Dialog.Description className="mb-5 mt-2.5 text-[15px] leading-normal text-mauve11">
                                                <Flex>
                                                    <Text>
                                                        {`Add "${data.title}" to blog...`}
                                                    </Text>
                                                    <Select.Root defaultValue="apple">
                                                        <Select.Trigger />
                                                        <Select.Content>
                                                            <Select.Group>
                                                                <Select.Item value="orange">
                                                                    This
                                                                </Select.Item>
                                                                <Select.Item value="apple">
                                                                    This is
                                                                    updated blog
                                                                </Select.Item>
                                                                <Select.Item value="grape">
                                                                    This another
                                                                </Select.Item>
                                                            </Select.Group>
                                                        </Select.Content>
                                                    </Select.Root>
                                                </Flex>
                                            </Dialog.Description>

                                            <div className="mt-[25px] flex justify-end">
                                                <Dialog.Close asChild>
                                                    <Button>Confirm</Button>
                                                </Dialog.Close>
                                            </div>
                                            <Dialog.Close asChild>
                                                <button
                                                    className="absolute right-2.5 top-2.5 inline-flex size-[25px] appearance-none items-center justify-center rounded-full text-violet11 hover:bg-violet4 focus:shadow-[0_0_0_2px] focus:shadow-violet7 focus:outline-none"
                                                    aria-label="Close"
                                                >
                                                    <Cross2Icon />
                                                </button>
                                            </Dialog.Close>
                                        </Dialog.Content>
                                    </Dialog.Portal>
                                </Dialog.Root>
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
