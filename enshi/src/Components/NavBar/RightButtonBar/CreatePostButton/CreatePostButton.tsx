import { PlusIcon } from "@radix-ui/react-icons";
import { Button, Text } from "@radix-ui/themes";
import { useTranslation } from "react-i18next";
import { Link } from "react-router-dom";

export default function CreatePostButton() {
    const { t } = useTranslation();

    return (
        <Link to={"/create"}>
            <Button variant="ghost" className="h-full">
                <PlusIcon />
                <Text>{t("createPost")}</Text>
            </Button>
        </Link>
    );
}
