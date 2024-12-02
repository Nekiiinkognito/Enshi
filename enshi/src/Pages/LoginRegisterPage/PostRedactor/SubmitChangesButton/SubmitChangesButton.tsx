import { Button } from "@radix-ui/themes";
import { useMutation } from "@tanstack/react-query";
import { useState } from "react";
import { useTranslation } from "react-i18next";
import { useNavigate, useParams } from "react-router-dom";
import { axiosLocalhost } from "../../../../api/axios/axios";

type TSubmitChangesButton = {
    className: string;
    titleValue: string;
    contentValue: string;
};

export default function SubmitChangesButton(props: TSubmitChangesButton) {
    const { t } = useTranslation();

    const [isDisabled, setIsDisabled] = useState(false);

    const navigate = useNavigate();
    const queryParams = useParams();

    const postMutation = useMutation({
        mutationFn: async () => {
            if (!props.titleValue) throw new Error("no title provided");
            if (!props.contentValue || props.contentValue === "<p><br></p>")
                throw new Error("no content provided");

            axiosLocalhost.put(`/posts/${queryParams["postId"]}`, {
                title: props.titleValue,
                content: props.contentValue,
            });
        },
        onMutate: () => {
            setIsDisabled(true);
        },
        onError: () => {
            setIsDisabled(false);
        },
        onSuccess: () => {
            navigate("/");
        },
    });

    return (
        <Button
            onClick={() => {
                postMutation.mutate();
            }}
            className={props.className}
            variant="soft"
            size={"4"}
            disabled={isDisabled}
        >
            {t("updatePost")}
        </Button>
    );
}
