import "./App.css";
import "@radix-ui/themes/styles.css";
import {
    Badge,
    Button,
    Callout,
    Card,
    Container,
    Flex,
    Text,
    TextArea,
    TextField,
    Theme,
    ThemePanel,
} from "@radix-ui/themes";

import { useTranslation } from "react-i18next";
import * as Form from "@radix-ui/react-form";

function App() {
    const { t } = useTranslation();

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

            <Container>
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
            </Container>

            <ThemePanel />
        </Theme>
    );
}

export default App;
