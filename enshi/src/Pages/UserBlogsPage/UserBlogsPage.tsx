import * as Dialog from "@radix-ui/react-dialog";
import { Cross2Icon, PlusIcon } from "@radix-ui/react-icons";
import {
    Box,
    Button,
    Container,
    Flex,
    Separator,
    Text,
} from "@radix-ui/themes";
import { useQuery } from "@tanstack/react-query";
import { axiosLocalhost } from "../../api/axios/axios";
import BlogBox from "../../Components/BlogBox/BlogBox";
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

                    <Dialog.Root>
                        <Dialog.Trigger asChild>
                            <Button onClick={() => {}}>
                                <PlusIcon />
                            </Button>
                        </Dialog.Trigger>
                        <Dialog.Portal>
                            <Dialog.Overlay className="fixed inset-0 bg-blackA6 data-[state=open]:animate-overlayShow" />
                            <Dialog.Content className="fixed left-1/2 top-1/2 max-h-[85vh] w-[90vw] max-w-[450px] -translate-x-1/2 -translate-y-1/2 rounded-md bg-white p-[25px] shadow-[hsl(206_22%_7%_/_35%)_0px_10px_38px_-10px,_hsl(206_22%_7%_/_20%)_0px_10px_20px_-15px] focus:outline-none data-[state=open]:animate-contentShow">
                                <Dialog.Title className="m-0 text-[17px] font-medium text-mauve12">
                                    Create blog
                                </Dialog.Title>
                                <Dialog.Description className="mb-5 mt-2.5 text-[15px] leading-normal text-mauve11">
                                    Create your new blog.
                                </Dialog.Description>
                                <fieldset className="mb-[15px] flex items-center gap-5">
                                    <label
                                        className="w-[90px] text-right text-[15px] text-violet11"
                                        htmlFor="title"
                                    >
                                       Blog title
                                    </label>
                                    <input
                                        className="inline-flex h-[35px] w-full flex-1 items-center justify-center rounded px-2.5 text-[15px] leading-none text-violet11 shadow-[0_0_0_1px] shadow-violet7 outline-none focus:shadow-[0_0_0_2px] focus:shadow-violet8"
                                        id="title"
                                        defaultValue="My blog"
                                    />
                                </fieldset>
                                <fieldset className="mb-[15px] flex items-center gap-5">
                                    <label
                                        className="w-[90px] text-right text-[15px] text-violet11"
                                        htmlFor="Description"
                                    >
                                        Description
                                    </label>
                                    <textarea
                                        className="pt-2 inline-flex h-[35px] w-full flex-1 items-center justify-center rounded px-2.5 text-[15px] leading-none text-violet11 shadow-[0_0_0_1px] shadow-violet7 outline-none focus:shadow-[0_0_0_2px] focus:shadow-violet8"
                                        id="Description"
                                        placeholder="Your description..."
                                    />
                                </fieldset>
                                <div className="mt-[25px] flex justify-end">
                                    <Dialog.Close asChild>
                                        <Button>
                                            Create blog
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
                            </Dialog.Content>
                        </Dialog.Portal>
                    </Dialog.Root>
                </Flex>
            </Container>
        </Box>
    );
}
