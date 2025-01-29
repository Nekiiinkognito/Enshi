import { Box, Skeleton } from "@radix-ui/themes";
import { useQuery } from "@tanstack/react-query";
import { axiosLocalhost } from "../../../api/axios/axios";

type TVoteCounter = {
    postId: string;
};

export default function VoteCounter(props: TVoteCounter) {
    const { data, isLoading } = useQuery({
        queryKey: [`post_vote_counter_${props.postId}` ],
        queryFn: async () => {
            const response = await axiosLocalhost.get(
                `post-votes/${props.postId}`
            );
            return response.data as { upvotes: number; downvotes: number };
        },
        gcTime: 1000 * 60, 
    });

    const calculateRating = (upvotes: number, downvotes: number) => {
        return upvotes + (-downvotes)
    }

    if (isLoading) {
        return <Skeleton className="w-4">
            {calculateRating(0, 0)}
        </Skeleton>
    }

    return <Box className="flex justify-center w-4 gap-2">
        {calculateRating(data?.upvotes || 0, data?.downvotes || 0)}
    </Box>;
}
