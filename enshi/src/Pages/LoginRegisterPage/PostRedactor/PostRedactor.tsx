import { Box, Container, Flex, Spinner } from "@radix-ui/themes";
import { useQuery } from "@tanstack/react-query";
import { useState } from "react";
import { useParams } from "react-router-dom";
import { axiosLocalhost } from "../../../api/axios/axios";
import Editor from "../../../Components/Editor/Editor";
import SubmitChangesButton from "./SubmitChangesButton/SubmitChangesButton";

export default function PostRedactor() {
    const [contentValue, setContentValue] = useState("");
    const [titleValue, setTitleValue] = useState("");

    const queryParams = useParams();

    const { isPending } = useQuery({
        queryKey: ["changePostKey", queryParams.postId],
        queryFn: async () => {
            try {
                const response = await axiosLocalhost.get(
                    `/posts/${queryParams.postId}`
                );

                setTitleValue(response.data["title"]);
                setContentValue(response.data["content"]);

                return response.data;
            } catch (error) {
                console.log(error);

                return error;
            }
        },
        gcTime: Infinity,
    });

    return (
        <>
            <Box className="flex flex-col flex-1">
                <Flex gap={"4"} direction={"column"} className="flex-[1]">
                    <Container className="flex-[1]">
                        <input
                            disabled={isPending}
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
                        {isPending ? (
                            <Spinner />
                        ) : (
                            <Editor
                                defaultValue={contentValue}
                                onChange={setContentValue}
                            />
                        )}
                    </Container>

                    <Box className="flex justify-center flex-[1] mb-4">
                        <SubmitChangesButton 
                        contentValue={contentValue}
                        titleValue={titleValue}
                        className="text-2xl rounded-full w-52" />
                    </Box>
                </Flex>
            </Box>
        </>
    );
}
