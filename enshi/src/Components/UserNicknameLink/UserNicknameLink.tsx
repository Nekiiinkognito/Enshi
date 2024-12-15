import { Skeleton, Text } from "@radix-ui/themes";
import { useQuery } from "@tanstack/react-query";
import { Link } from "react-router-dom";
import { axiosLocalhost } from "../../api/axios/axios";

type TUserNicknameLink = {
    userId: string;
};

export default function UserNicknameLink(props: TUserNicknameLink) {
    const { data, isPending } = useQuery({
        queryKey: [`userLink${props.userId}`],
        queryFn: async () => {
            const response = await axiosLocalhost.get(
                `/user/${props.userId || 0}`
            );
            return response.data as string;
        },
    });

    if (isPending)
        return (
            <Skeleton>
                <Text>@Nickname</Text>
            </Skeleton>
        );

    return (
        <Link to={`/users/${data}`}>
            <Text>@{data}</Text>
        </Link>
    );
}
