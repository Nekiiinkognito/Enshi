import { Button } from "@radix-ui/themes";
import { useMutation } from "@tanstack/react-query";
import { useAtom } from "jotai";
import { useState } from "react";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router-dom";
import { axiosLocalhost } from "../../../api/axios/axios";
import {
    postCreationAtom,
    postCreationTitleAtom,
} from "../../../AtomStore/AtomStore";

type TSubmitPostButton = {
    className: string;
};

export default function SubmitPostButton(props: TSubmitPostButton) {
    const { t } = useTranslation();

    const [isDisabled, setIsDisabled] = useState(false);

    const [contentValue, setContentValue] = useAtom(postCreationAtom);
    const [titleValue, setTitleValue] = useAtom(postCreationTitleAtom);

    const navigate = useNavigate();

    const postMutation = useMutation({
        mutationFn: async () => {
            if (!titleValue) throw new Error("no title provided");
            if (!contentValue || contentValue === "<p><br></p>")
                throw new Error("no content provided");

            axiosLocalhost.post("/posts", {
                title: titleValue,
                content: contentValue,
            });
        },
        onMutate: () => {
            setIsDisabled(true);
        },
        onError: () => {
            setIsDisabled(false);
        },
        onSuccess: () => {
            setContentValue("");
            setTitleValue("");
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
            {t("submit")}
        </Button>
    );
}
