import * as Form from "@radix-ui/react-form";
import { CrossCircledIcon } from "@radix-ui/react-icons";
import { Button, Card, Heading, Text, TextField } from "@radix-ui/themes";
import { useEffect, useState } from "react";
import ShowPasswordButton from "../ShowPasswordButton/ShowPasswordButton";

export default function LoginElement() {
    const [showPassword, setShowPassword] = useState(false);
    const [isCapsLockOn, setIsCapsLockOn] = useState(false);

    useEffect(() => {
        const f = (e: KeyboardEvent) => {
            if (e.getModifierState("CapsLock")) {
                setIsCapsLockOn(true);
            } else {
                setIsCapsLockOn(false);
            }
        };

        document.addEventListener("keydown", f);

        return () => {
            document.removeEventListener("keydown", f);
        };
    }, []);

    return (
        <Card
            size={"2"}
            className="absolute w-1/5 
                        left-[50%] top-[50%] 
                        translate-x-[-50%] translate-y-[-50%]"
        >
            <Heading weight={"medium"} className="mb-4 text-center">
                Log in form
            </Heading>
            <Form.Root>
                <Form.Field className="mb-2.5 gap-0.5 grid" name="username">
                    <div className="flex items-baseline justify-between gap-2">
                        <Form.Label>
                            <Heading size={"3"}>Username</Heading>
                        </Form.Label>
                        <Form.Message match="valueMissing">
                            <Text color="red">Please enter your username</Text>
                        </Form.Message>
                    </div>
                    <Form.Control asChild>
                        <TextField.Root type="text" required>
                            <Form.ValidityState>
                                {(validity) => (
                                    <TextField.Slot
                                        side="right"
                                        color="red"
                                        className={
                                            validity
                                                ? validity.valid
                                                    ? "hidden"
                                                    : "mr-0.5"
                                                : "hidden"
                                        }
                                    >
                                        <CrossCircledIcon />
                                    </TextField.Slot>
                                )}
                            </Form.ValidityState>
                        </TextField.Root>
                    </Form.Control>
                </Form.Field>

                <Form.Field className="mb-2.5 gap-0.5 grid" name="password">
                    <div className="flex items-baseline justify-between gap-2">
                        <Form.Label>
                            <Heading size={"3"}>Password</Heading>
                        </Form.Label>
                        <Form.Message match="valueMissing">
                            <Text color="red">Please enter your password</Text>
                        </Form.Message>
                    </div>
                    <Form.Control asChild>
                        <TextField.Root
                            type={showPassword ? "text" : "password"}
                            required
                        >
                            <Form.ValidityState>
                                {(validity) => (
                                    <TextField.Slot
                                        side="right"
                                        color={
                                            validity
                                                ? validity.valid
                                                    ? undefined
                                                    : "red"
                                                : undefined
                                        }
                                    >
                                        <ShowPasswordButton
                                            isShown={showPassword}
                                            setIsShown={setShowPassword}
                                        />
                                    </TextField.Slot>
                                )}
                            </Form.ValidityState>
                        </TextField.Root>
                    </Form.Control>
                    <Text size={"1"} hidden={!isCapsLockOn}>
                        Caps lock is on
                    </Text>
                </Form.Field>

                {/* <Form.Field
                    className="mb-2.5 gap-0.5 grid"
                    name="conf-password"
                >
                    <div className="flex items-baseline justify-between">
                        <Form.Label>
                            <Heading size={"3"}>Confirm password</Heading>
                        </Form.Label>
                        <Form.Message match="valueMissing">
                            <Text color="red">Please enter your password</Text>
                        </Form.Message>
                        <Form.Message
                            match={(value, formData) =>
                                value !== formData.get("password")
                            }
                        >
                            <Text color="red">Passwords must be the same</Text>
                        </Form.Message>
                    </div>
                    <Form.Control asChild>
                        <TextField.Root
                            type={showConfPassword ? "text" : "password"}
                            required
                        >
                            <Form.ValidityState>
                                {(validity) => (
                                    <TextField.Slot
                                        side="right"
                                        color={
                                            validity
                                                ? validity.valid
                                                    ? undefined
                                                    : "red"
                                                : undefined
                                        }
                                    >
                                        <ShowPasswordButton
                                            isShown={showConfPassword}
                                            setIsShown={setShowConfPassword}
                                        />
                                    </TextField.Slot>
                                )}
                            </Form.ValidityState>
                        </TextField.Root>
                    </Form.Control>
                </Form.Field> */}

                <Form.Submit className="flex justify-center w-full">
                    <Button type="submit" className="w-1/3">
                        <Text size={"3"}>Submit</Text>
                    </Button>
                </Form.Submit>
            </Form.Root>
        </Card>
    );
}
