import * as Dialog from "@radix-ui/react-dialog";
import { Cross2Icon, PlusIcon } from "@radix-ui/react-icons";
import { Box, Button, Heading, Text, Theme } from "@radix-ui/themes";
import { useMutation } from "@tanstack/react-query";
import { useState } from "react";
import { axiosLocalhost } from "../../../api/axios/axios";
import queryClient from "../../../api/QueryClient/QueryClient";
import useToast from "../../../hooks/useToast";

export default function BlogCreationDialog() {
    const createToast = useToast();

    const [title, setTitle] = useState<string>("My blog");
    const [description, setDescription] = useState<string>("");

    const addMutation = useMutation({
        onMutate: () => {},
        mutationFn: async () => {
            await axiosLocalhost.post("/blogs", {
                title,
                description,
            });
        },
        onSuccess: () => {
            createToast({
                title: `Success!`,
                description: `Blog created successfully!`,
            });
        },
        onError: (_error) => {
            createToast({
                title: `Error!`,
                description: `Blog creation failed!`,
            });
        },
        onSettled: () => {
            queryClient.invalidateQueries({
                queryKey: ["userBlogs"],
            });
        },
    });

    return (
        <Box>
            <Dialog.Root>
                <Dialog.Trigger asChild>
                    <Button onClick={() => {}} className="w-full">
                        <PlusIcon />
                    </Button>
                </Dialog.Trigger>
                <Dialog.Portal>
                    <Box>
                        <Theme>
                            <Dialog.Overlay className="fixed inset-0 bg-black/40 backdrop-blur-sm data-[state=open]:animate-appear" />
                            <Dialog.Content className="fixed left-1/2 top-1/2 max-h-[85vh] w-[90vw] max-w-[600px] -translate-x-1/2 -translate-y-1/2 rounded-md bg-white p-[25px] shadow-[hsl(206_22%_7%_/_35%)_0px_10px_38px_-10px,_hsl(206_22%_7%_/_20%)_0px_10px_20px_-15px] focus:outline-none data-[state=open]:animate-appear">
                                <Dialog.Title className="m-0 text-[17px] font-medium text-mauve12">
                                    <Heading>Create blog</Heading>
                                </Dialog.Title>
                                <Dialog.Description className="mb-5 mt-2.5 text-[15px] leading-normal text-mauve11">
                                    <Text>Create your new blog.</Text>
                                </Dialog.Description>
                                <fieldset className="mb-[15px] flex items-center gap-5">
                                    <label
                                        className="w-[90px] text-right text-[15px] text-violet11"
                                        htmlFor="title"
                                    >
                                        <Text>Blog title</Text>
                                    </label>
                                    <input
                                        className="inline-flex h-[35px] w-full flex-1 items-center justify-center rounded px-2.5 text-[15px] leading-none text-violet11 shadow-[0_0_0_1px] shadow-violet7 outline-none focus:shadow-[0_0_0_2px] focus:shadow-violet8"
                                        id="title"
                                        defaultValue="My blog"
                                        onChange={(e) => {
                                            setTitle(e.target.value);
                                        }}
                                    />
                                </fieldset>
                                <fieldset className="mb-[15px] flex items-center gap-5">
                                    <label
                                        className="w-[90px] text-right text-[15px] text-violet11"
                                        htmlFor="Description"
                                    >
                                        <Text>Description</Text>
                                    </label>
                                    <textarea
                                        className="pt-2 inline-flex h-[35px] w-full flex-1 items-center justify-center rounded px-2.5 text-[15px] leading-none text-violet11 shadow-[0_0_0_1px] shadow-violet7 outline-none focus:shadow-[0_0_0_2px] focus:shadow-violet8"
                                        id="Description"
                                        placeholder="Your description..."
                                        onChange={(e) => {
                                            setDescription(e.target.value);
                                        }}
                                    />
                                </fieldset>
                                <div className="mt-[25px] flex justify-end">
                                    <Dialog.Close
                                        asChild
                                        onClick={() => {
                                            addMutation.mutate();
                                        }}
                                    >
                                        <Button>Create blog</Button>
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
                        </Theme>
                    </Box>
                </Dialog.Portal>
            </Dialog.Root>
        </Box>
    );
}
