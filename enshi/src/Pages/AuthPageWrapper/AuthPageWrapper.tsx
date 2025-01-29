import { Container, Text } from "@radix-ui/themes";
import { t } from "i18next";
import { useAtomValue } from "jotai";
import React from "react";
import { useNavigate } from "react-router-dom";
import { userAtom } from "../../AtomStore/AtomStore";

export default function AuthPageWrapper(props: React.PropsWithChildren) {
    const user = useAtomValue(userAtom);
    const navigate = useNavigate();

    if (!user) {
        navigate("/login");
        return (
            <Container size={"4"} className="mt-4">
                <Text size={"7"}>{t("errors.unauthorized")}</Text>
            </Container>
        );
    }

    return props.children;
}
