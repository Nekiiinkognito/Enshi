import "./App.css";
import "@radix-ui/themes/styles.css";
import {
    Badge,
    Box,
    Button,
    Callout,
    Container,
    Flex,
    Separator,
    Text,
    Theme,
    ThemePanel,
} from "@radix-ui/themes";

import { useTranslation } from "react-i18next";
import { useEffect, useRef, useState } from "react";
import parse from "html-react-parser";
import Quill from "quill/core";
import Editor from "./Components/Editor/Editor";

const Delta = Quill.import("delta");

function App() {
    const { t } = useTranslation();
    const [value, setValue] = useState<string>();
    const first = useRef(null);

    return (
        <Theme className="h-fit" accentColor="indigo" grayColor="slate">
            <Flex className="w-full">
                <Callout.Root>
                    <Text>Hello world!</Text>
                    <Button highContrast variant="soft">
                        {t("hello")}
                    </Button>
                    <Button variant="soft">Let's goooo!</Button>
                    <Badge color="red">What</Badge>
                </Callout.Root>
            </Flex>

            {/* <Container>
                <Card className="w-full top-10">
                    <Form.Root className="FormRoot">
                        <Form.Field className="FormField" name="email">
                            <div
                                style={{
                                    display: "flex",
                                    alignItems: "baseline",
                                    justifyContent: "space-between",
                                }}
                            >
                                <Form.Label className="FormLabel">
                                    Email
                                </Form.Label>
                                <Form.Message
                                    className="FormMessage"
                                    match="valueMissing"
                                >
                                    Please enter your email
                                </Form.Message>
                                <Form.Message
                                    className="FormMessage"
                                    match="typeMismatch"
                                >
                                    Please provide a valid email
                                </Form.Message>
                            </div>
                            <Form.Control asChild>
                                <TextField.Root
                                    className="Input"
                                    type="email"
                                    required
                                />
                            </Form.Control>
                        </Form.Field>
                        <Form.Field className="FormField" name="question">
                            <div
                                style={{
                                    display: "flex",
                                    alignItems: "baseline",
                                    justifyContent: "space-between",
                                }}
                            >
                                <Form.Label className="FormLabel">
                                    Question
                                </Form.Label>
                                <Form.Message
                                    className="FormMessage"
                                    match="valueMissing"
                                >
                                    Please enter a question
                                </Form.Message>
                            </div>
                            <Form.Control asChild>
                                <TextArea className="Textarea" required />
                            </Form.Control>
                        </Form.Field>
                        <Form.Submit asChild>
                            <Button variant="soft" style={{ marginTop: 10 }}>
                                Post question
                            </Button>
                        </Form.Submit>
                    </Form.Root>
                </Card>
            </Container> */}
            <>
                <Container>
                    <Editor />
                </Container>
            </>

            <>
                <Container className="mt-4">
                    <Separator orientation={"horizontal"} className="w-full" />
                    {parse(value || "")}
                </Container>
            </>
            <ThemePanel />
        </Theme>
    );
}

export default App;
