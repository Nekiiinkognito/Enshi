import * as Dialog from "@radix-ui/react-dialog";
import { Cross2Icon } from "@radix-ui/react-icons";
import { Button, Card, Flex, Select, Text, Theme } from "@radix-ui/themes";
import { useQuery } from "@tanstack/react-query";
import { useState } from "react";
import { Blog } from "../../../@types/BlogTypes";
import { axiosLocalhost } from "../../../api/axios/axios";
import { JSONWithInt64 } from "../../../utils/idnex";

export default function AddPostToBlogDialog() {
    const { data } = useQuery({
        queryKey: ["userBlogs"],
        queryFn: async () => {
            const response = await axiosLocalhost.get("/user/blogs", {
                transformResponse: [(data) => data],
            });

            let temp = JSONWithInt64(response.data);

            return temp as Blog[];
        },
    });

    const [selectedBlog, setSelectedBlog] = useState<string>("");

    return (
        <Dialog.Root>
            <Dialog.Trigger asChild>
                <Button variant="surface" className="h-5">
                    Add to blog
                </Button>
            </Dialog.Trigger>
            <Dialog.Portal>
                <Dialog.Overlay className="fixed inset-0 bg-black/40 data-[state=open]:animate-overlayShow" />
                <Dialog.Content className="fixed left-1/2 top-1/2 max-h-[85vh] w-[90vw] max-w-[600px] -translate-x-1/2 -translate-y-1/2 rounded-md p-[25px] shadow-[hsl(206_22%_7%_/_35%)_0px_10px_38px_-10px,_hsl(206_22%_7%_/_20%)_0px_10px_20px_-15px] focus:outline-none data-[state=open]:animate-contentShow">
                    <Theme>
                        <Card>
                            <Dialog.Title className="m-0 text-[17px] font-medium text-mauve12">
                                Add this post to blog
                            </Dialog.Title>
                            <Dialog.Description className="mb-5 mt-2.5 text-[15px] leading-normal text-mauve11">
                                <Flex gap={"2"} align={"center"}>
                                    <Text>{`Add post to `}</Text>
                                    <Select.Root>
                                        <Select.Trigger className="w-40"/>
                                        <Select.Content>
                                            <Select.Group>
                                                {data?.map((blog, i) => {
                                                    return (
                                                        <Select.Item key={i} value={`${blog.blog_id}`}>
                                                            {blog.title}
                                                        </Select.Item>
                                                    );
                                                })}
                                            </Select.Group>
                                        </Select.Content>
                                    </Select.Root>
                                </Flex>
                            </Dialog.Description>

                            <div className="mt-[25px] flex justify-end">
                                <Dialog.Close asChild>
                                    <Button variant="outline" color="green">
                                        Confirm
                                    </Button>
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
                        </Card>
                    </Theme>
                </Dialog.Content>
            </Dialog.Portal>
        </Dialog.Root>
    );
}
