import { Container } from "@radix-ui/themes";
import { useState } from "react";
import { useTranslation } from "react-i18next";
import ArticleViewer from "../../Components/ArticleViewer/ArticleViewer";
import Editor from "../../Components/Editor/Editor";

export default function PostCreatorPage() {
    const [userInput, setUserInput] = useState("");

    const { t } = useTranslation();

    return (
        <>
            <Container className="mt-10">
                <input
                    placeholder={"Post title"}
                    className="mb-2 border-0 border-b-[1px] 
                                outline-none w-full border-b-gray-400
                                text-[60px] pl-4 pr-4 font-times"
                ></input>
                <Editor onChange={setUserInput} />
                <ArticleViewer htmlToParse={userInput} />
            </Container>
        </>
    );
}
