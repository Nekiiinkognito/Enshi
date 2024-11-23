import { Container } from "@radix-ui/themes";
import { Interweave } from "interweave";
import { useParams } from "react-router-dom";

type TArticleViewer = {
    htmlToParse?: string;
}

export default function ArticleViewer(props: TArticleViewer) {
    const queryPapms = useParams()

    return (
        <>
            <div className="ql-snow">
                <Container className="mt-4 ql-editor">
                    <Interweave content={props?.htmlToParse || ""} /> 
                </Container>
            </div>
        </>
    );
}
