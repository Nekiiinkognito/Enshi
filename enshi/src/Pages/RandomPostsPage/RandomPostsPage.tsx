import * as ScrollArea from "@radix-ui/react-scroll-area";
import { Container, Flex, Heading, Separator } from "@radix-ui/themes";
import { useQuery } from "@tanstack/react-query";
import { useTranslation } from "react-i18next";
import { GetRandomPostsRow } from "../../@types/PostTypes";
import { axiosLocalhost } from "../../api/axios/axios";
import PostCard from "./PostCard/PostCard";

const LIMIT = 10;

export default function RandomPostsPage() {
    const {t} = useTranslation()

    const { data, refetch } = useQuery({
        queryKey: ["random_posts_key"],
        queryFn: async () => {
            try {
                const response = await axiosLocalhost.get(
                    `/posts/random?limit=${LIMIT}`
                );

                return response.data as GetRandomPostsRow[];
            } catch (error) {
                console.log(`Something went wrong`);
            }

            return [];
        },
    });

    return (
        <>
                <Flex direction={"column"} className="mx-auto">
                    <Heading size={"9"} weight={"regular"} className="text-center">
                        {t("discover")}
                    </Heading>

                    <Separator size={"4"} className="my-8" />

                    <ScrollArea.Root className="w-full h-full overflow-hidden">
                        <ScrollArea.Viewport className="overflow-scroll rounded size-full">
                            {data?.map((post, i) => {
                                return (
                                    <Container size={"3"} key={`post${i}`}>
                                        <PostCard post={post} />
                                    </Container>
                                );
                            })}
                        </ScrollArea.Viewport>
                        <ScrollArea.Scrollbar
                            className="z-50 flex touch-none select-none p-0.5 w-2"
                            orientation="vertical"
                        >
                            <ScrollArea.Thumb className="relative flex-1 rounded-[10px] bg-slate-200"/>
                        </ScrollArea.Scrollbar>
                        {/* <ScrollArea.Scrollbar
                            className="flex touch-none select-none bg-blackA3 p-0.5 transition-colors duration-[160ms] ease-out hover:bg-blackA5 data-[orientation=horizontal]:h-2.5 data-[orientation=vertical]:w-2.5 data-[orientation=horizontal]:flex-col"
                            orientation="horizontal"
                        >
                            <ScrollArea.Thumb className="relative flex-1 rounded-[10px] bg-mauve10 before:absolute before:left-1/2 before:top-1/2 before:size-full before:min-h-[44px] before:min-w-[44px] before:-translate-x-1/2 before:-translate-y-1/2" />
                        </ScrollArea.Scrollbar> */}
                        {/* <ScrollArea.Corner className="bg-blackA5" /> */}
                    </ScrollArea.Root>
                </Flex>
        </>
    );
}
