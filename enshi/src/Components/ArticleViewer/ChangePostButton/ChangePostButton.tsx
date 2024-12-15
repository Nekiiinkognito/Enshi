import { Button } from "@radix-ui/themes";
import { useNavigate } from "react-router-dom";

type TChangePostButton = {
    postId: number | string;
};

export default function ChangePostButton(props: TChangePostButton) {
    const navigate = useNavigate();

    return (
        <Button
            size={"1"}
            className="h-5"
            variant="surface"
            onClick={() => navigate("/posts/change/" + props.postId)}
        >
            {"Change article"}
        </Button>
    );
}
