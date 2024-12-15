import { Box, Container, Flex } from "@radix-ui/themes";
import { useAtom, useSetAtom } from "jotai";
import {
    postCreationAtom,
    postCreationTitleAtom
} from "../../AtomStore/AtomStore";
import Editor from "../../Components/Editor/Editor";
import SubmitPostButton from "./SubmitPostButton/SubmitPostButton";

export default function PostCreatorPage() {
    const [titleValue, setTitleValue] = useAtom(postCreationTitleAtom);
    const setContentValue = useSetAtom(postCreationAtom);
    
    return (
        <>
            
            <Box className="flex flex-col flex-1">
                <Flex gap={"4"} direction={"column"} className="flex-[1]">
                    <Container className="flex-[1]">
                        <input
                            placeholder={"Post title"}
                            className="mb-2 border-0 border-b-[1px] 
                                outline-none w-full border-b-gray-400
                                text-[60px] pl-4 pr-4 font-times"
                            onChange={(e) => {
                                setTitleValue(e.target.value);
                            }}
                            value={titleValue}
                        />
                    </Container>

                    <Container className="overflow-y-auto flex-grow-[100]">
                        <Editor onChange={setContentValue} />
                    </Container>

                    <Box className="flex justify-center flex-[1] mb-4">
                        <SubmitPostButton className="text-2xl rounded-full w-52" />
                    </Box>
                </Flex>
            </Box>
        </>
    );
}
