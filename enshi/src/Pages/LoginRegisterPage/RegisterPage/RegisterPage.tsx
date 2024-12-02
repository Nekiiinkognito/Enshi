import * as Form from "@radix-ui/react-form";
import { CrossCircledIcon } from "@radix-ui/react-icons";
import { Button, Card, Heading, Text, TextField } from "@radix-ui/themes";
import { useMutation } from "@tanstack/react-query";
import { useSetAtom } from "jotai";
import { useState } from "react";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router-dom";
import { axiosLocalhost } from "../../../api/axios/axios";
import { userAtom } from "../../../AtomStore/AtomStore";
import UseCapsLock from "../../../hooks/useCapsLock";
import ShowPasswordButton from "../ShowPasswordButton/ShowPasswordButton";

type TRegisterData = {
    username: string;
    password: string;
    email: string;
};

export default function RegisterPage() {
    const setUserAtom = useSetAtom(userAtom)
    const [showPassword, setShowPassword] = useState(false);
    const [showConfPassword, setShowConfPassword] = useState(false);
    const { isCapsLockOn } = UseCapsLock();

    const { t } = useTranslation();

    const [isError, setIsError] = useState(false);

    const navigate = useNavigate();

    const registerMutation = useMutation({
        mutationFn: async (data: TRegisterData) => {
            let response = await axiosLocalhost.post("/users", JSON.stringify(data));
            setUserAtom({
                username: response.data.username,
                isAdmin: false,
                id: response.data.id,
            })
        },

        onError: (error, _variables, _context) => {
            console.log(error);
            setIsError(true);
        },

        onSuccess: () => {
            navigate("/");
        },
    });

    return (
        <Card
            size={"2"}
            className="absolute w-[25rem] min-w-[20rem]
                        left-[50%] top-[50%] 
                        translate-x-[-50%] translate-y-[-50%]"
        >
            <Heading weight={"medium"} className="mb-4 text-center">
                {t("registerForm")}
            </Heading>
            <Form.Root
            className="flex flex-col gap-2"
                onSubmit={(e) => {
                    e.preventDefault();
                    let formData = new FormData(
                        document.querySelector("form") as HTMLFormElement
                    );

                    let registerData: TRegisterData = {
                        password: (formData.get("password") as string) || "",
                        username: (formData.get("username") as string) || "",
                        email: (formData.get("email") as string) || "",
                    };

                    registerMutation.mutate(registerData);
                }}
            >
                <Form.Field className="gap-0.5 grid" name="username">
                    <div className="flex items-baseline justify-between gap-2">
                        <Form.Label>
                            <Text size={"4"}>{t("username")}</Text>
                        </Form.Label>
                        <Form.Message match="valueMissing">
                            <Text color="red">{t("errors.enterUsername")}</Text>
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

                <Form.Field className="gap-0.5 grid" name="email">
                    <div className="flex items-baseline justify-between gap-2">
                        <Form.Label>
                            <Text size={"4"}>{t("email")}</Text>
                        </Form.Label>
                        <Form.Message match="valueMissing">
                            <Text color="red">{t("errors.enterEmail")}</Text>
                        </Form.Message>
                        <Form.Message match="typeMismatch">
                            <Text color="red">{t("errors.invalidEmail")}</Text>
                        </Form.Message>
                    </div>
                    <Form.Control asChild>
                        <TextField.Root type="email" required>
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

                <Form.Field className="gap-0.5 grid" name="password">
                    <div className="flex items-baseline justify-between gap-2">
                        <Form.Label>
                            <Text size={"4"}>{t("password")}</Text>
                        </Form.Label>
                        <Form.Message match="valueMissing">
                            <Text color="red">{t("errors.enterPassword")}</Text>
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
                        {t("capsLogWarning")}
                    </Text>
                </Form.Field>

                <Form.Field
                    className="gap-0.5 grid"
                    name="conf-password"
                >
                    <div className="flex items-baseline justify-between">
                        <Form.Label>
                            <Text size={"4"}>{t("confirmPassword")}</Text>
                        </Form.Label>
                        <Form.Message match="valueMissing">
                            <Text color="red">{t("errors.enterPassword")}</Text>
                        </Form.Message>
                        <Form.Message
                            match={(value, formData) =>
                                value !== formData.get("password")
                            }
                        >
                            <Text color="red">
                                {t("errors.passwordsMismatch")}
                            </Text>
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
                    <Text size={"1"} hidden={!isCapsLockOn}>
                        {t("capsLogWarning")}
                    </Text>
                </Form.Field>

                <Text color="red" hidden={!isError}>
                    {t("errors.invalidRegisterData")}
                </Text>

                <Form.Submit className="flex justify-center mt-2" asChild>
                    <Button type="submit" className="w-full m-auto">
                        <Text size={"3"}>{t("submit")}</Text>
                    </Button>
                </Form.Submit>
            </Form.Root>
        </Card>
    );
}
