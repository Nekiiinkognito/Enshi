import { Button } from "@radix-ui/themes";
import { useMutation } from "@tanstack/react-query";
import { useAtomValue } from "jotai";
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
    const contentValue = useAtomValue(postCreationAtom);
    const titleValue = useAtomValue(postCreationTitleAtom);

    const navigate = useNavigate();

    const [isDisabled, setIsDisabled] = useState(false);

    const postMutation = useMutation({
        mutationFn: async () => {
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
