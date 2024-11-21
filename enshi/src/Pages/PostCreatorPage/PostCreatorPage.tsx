import { Container, Text } from "@radix-ui/themes";
import { useAtomValue } from "jotai";
import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import { userAtom } from "../../AtomStore/AtomStore";
import Editor from "../../Components/Editor/Editor";

export default function PostCreatorPage() {
    const user = useAtomValue(userAtom);

    const [userInput, setUserInput] = useState("");

    const { t } = useTranslation();

    useEffect(() => {
        console.log("asdasd", userInput);
        
    }, [userInput])
    

    if (!user) {
        return (
            <Container size={"4"} className="mt-4">
                <Text size={"7"}>{t("errors.unauthorized")}</Text>
            </Container>
        );
    }

    return (
        <>
            <Container className="mt-10">
                <Editor onChange={setUserInput} />
            </Container>
        </>
    );
}
