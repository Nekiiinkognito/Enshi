import { Box, Skeleton } from "@radix-ui/themes";
import { useQuery } from "@tanstack/react-query";
import { axiosLocalhost } from "../../../api/axios/axios";

type TVoteCounter = {
    postId: string;
};

export default function VoteCounter(props: TVoteCounter) {
    const { data, isLoading } = useQuery({
        queryKey: ["post_vote_counter"],
        queryFn: async () => {
            const response = await axiosLocalhost.get(
                `post-votes/${props.postId}`
            );
            return response.data as { upvotes: number; downvotes: number };
        },
    });

    const calculateRating = (upvotes: number, downvotes: number) => {
        return upvotes + (-downvotes)
    }

    if (isLoading) {
        return <Skeleton>
            {calculateRating(0, 0)}
        </Skeleton>
    }

    return <Box>
        {calculateRating(data?.upvotes || 0, data?.downvotes || 0)}
    </Box>;
}
