import { Container, Separator, Text } from "@radix-ui/themes";
import { useQuery } from "@tanstack/react-query";
import { Interweave } from "interweave";
import { useParams } from "react-router-dom";
import { axiosLocalhost } from "../../api/axios/axios";
import SkeletonLoader from "./SkeletonLoader/SkeletonLoader";

type TArticleViewer = {
    htmlToParse?: string;
};

export default function ArticleViewer(props: TArticleViewer) {
    let queryParams = useParams();

    const { data, isPending } = useQuery({
        queryKey: [`post_${queryParams["postId"]}`],
        queryFn: async () => {
            const response = await axiosLocalhost.get(
                `posts/${queryParams["postId"]}`
            );

            return response.data;
        },
    });

    if (isPending)
        return (
            <SkeletonLoader />
        );

    return (
        <>
            <div className="ql-snow">
                <Container size={"2"} className="mt-4">
                    <Text className="mb-2" as="div" size={"9"}>{data.title}</Text>
                    <Separator size={"4"} className="mb-2" />
                    <Interweave content={data.content} />
                </Container>
            </div>
        </>
    );
}
