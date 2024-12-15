import { DoubleArrowDownIcon, DoubleArrowUpIcon } from "@radix-ui/react-icons";
import { IconButton } from "@radix-ui/themes";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { axiosLocalhost } from "../../../api/axios/axios";

export const UPVOTE = true;
export const DOWNVOTE = false;

type TVoteButton = {
    postId: string;
    vote: boolean;
};

export default function VoteButton(props: TVoteButton) {
    const queryClient = useQueryClient();

    const { data } = useQuery({
        queryKey: [props.vote + "voteCheck"],
        queryFn: async () => {
            const response = await axiosLocalhost.get(
                `post-vote/${props.postId}`
            );

            return (response.data?.vote as boolean) === props.vote || false;
        },
        gcTime: 0,
    });

    const voteMutation = useMutation({
        mutationKey: [`voteMutation${props.vote}`],
        onMutate: async () => {
            queryClient.cancelQueries({ queryKey: [props.vote + "voteCheck"] });

            queryClient.setQueryData([props.vote + "voteCheck"], true);
            queryClient.setQueryData([!props.vote + "voteCheck"], false);
        },
        mutationFn: async () => {
            await axiosLocalhost.post(`post-votes/${props.postId}`, {
                vote: props.vote,
            });
        },
        onSuccess: () => {},
        onError: () => {
            queryClient.setQueryData([props.vote + "voteCheck"], false);
        },
        onSettled: () => {
            queryClient.invalidateQueries({
                queryKey: [props.vote + "voteCheck"],
            });
            queryClient.invalidateQueries({
                queryKey: ["post_vote_counter"],
            });
        },
    });

    return (
        <IconButton
            variant={data ? "solid" : "outline"}
            size={"1"}
            onClick={() => voteMutation.mutate()}
        >
            {props.vote ? <DoubleArrowUpIcon /> : <DoubleArrowDownIcon />}
        </IconButton>
    );
}
