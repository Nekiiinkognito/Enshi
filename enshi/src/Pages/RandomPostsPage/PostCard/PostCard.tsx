import { ImageIcon } from "@radix-ui/react-icons";
import { Box, Card, Heading } from "@radix-ui/themes";
import { useNavigate } from "react-router-dom";
import { GetRandomPostsRow } from "../../../@types/PostTypes";

type TPostCard = {
    post: GetRandomPostsRow;
};

export default function PostCard({ post }: TPostCard) {
    const navigate = useNavigate()

    const clickHandler = () => {
        navigate(`/posts/${post.post_id.toString()}`)
    }

    return (
        <Card className="h-32 mb-4" onClick={clickHandler}>
            <Box className="flex size-full">
                <Box>
                    <ImageIcon className="w-full h-full" />
                </Box>

                <Box className="px-4 pt-2">
                    <Heading>{post.title}</Heading>
                </Box>
            </Box>
        </Card>
    );
}
