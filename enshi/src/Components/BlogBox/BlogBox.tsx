import { Avatar, Card, Flex, Heading } from "@radix-ui/themes";
import { useNavigate } from "react-router-dom";
import UserNicknameLink from "../UserNicknameLink/UserNicknameLink";

type TBlogBox = {
    title?: string;
    blogId?: string;
    userId: string;
};

export default function BlogBox(props: TBlogBox) {
    const navigate = useNavigate();

    return (
        <Card className="w-full h-20" onClick={() => navigate(``)}>
            <Flex direction={"column"}>
                <Heading size={"4"}>{props?.title || "...No title..."}</Heading>
                <Flex align={"center"} gap={"2"} mt={"1"}>
                  <Avatar size={"2"} className="rounded-full" fallback={"SI"} />
                  <UserNicknameLink userId={props.userId} />
                </Flex>
            </Flex>
        </Card>
    );
}
