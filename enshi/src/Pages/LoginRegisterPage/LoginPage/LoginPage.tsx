import * as Form from "@radix-ui/react-form";
import { CrossCircledIcon } from "@radix-ui/react-icons";
import { Button, Card, Heading, Text, TextField } from "@radix-ui/themes";
import { useMutation } from "@tanstack/react-query";
import { t } from "i18next";
import { useAtom } from "jotai";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { axiosLocalhost } from "../../../api/axios/axios";
import { userAtom } from "../../../AtomStore/AtomStore";
import UseCapsLock from "../../../hooks/useCapsLock";
import ShowPasswordButton from "../ShowPasswordButton/ShowPasswordButton";

type TLoginData = {
    username: string;
    password: string;
};

export default function LoginPage() {
    const [userAtomValue, setUserAtom] = useAtom(userAtom)
    const [showPassword, setShowPassword] = useState(false);
    const { isCapsLockOn } = UseCapsLock();
    const [isError, setIsError] = useState(false);

    const navigate = useNavigate();

    const logInMutation = useMutation({
        mutationFn: async (data: TLoginData) => {
            let response = await axiosLocalhost.post("/login", JSON.stringify(data));
            setUserAtom({
                username: response.data.username,
                isAdmin: false
            })
        },

        onError: (error, _variables, _context) => {
            console.log(error);
            setIsError(true);
        },

        onSuccess: () => {
            let isAdminFunc = async () => {
                let response = await axiosLocalhost.get("/admin/check");
                if (response.status === 200) {
                    setUserAtom({
                        username: userAtomValue?.username || "",
                        isAdmin: true
                    })
                }
            };

            isAdminFunc();

            navigate("/");
        },
    });

    return (
        <Card
            size={"2"}
            className="absolute w-1/5 
                        left-[50%] top-[50%] 
                        translate-x-[-50%] translate-y-[-50%]"
        >
            <Heading weight={"medium"} className="mb-4 text-center">
                {t("loginForm")}
            </Heading>
            <Form.Root
                onSubmit={(e) => {
                    e.preventDefault();
                    let formData = new FormData(
                        document.querySelector("form") as HTMLFormElement
                    );

                    let loginData: TLoginData = {
                        password: (formData.get("password") as string) || "",
                        username: (formData.get("username") as string) || "",
                    };

                    logInMutation.mutate(loginData);
                }}
            >
                <Form.Field className="mb-2.5 gap-0.5 grid" name="username">
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

                <Form.Field className="mb-2.5 gap-0.5 grid" name="password">
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

                <Text color="red" hidden={!isError}>
                    {t("errors.invalidLoginData")}
                </Text>

                <Form.Submit className="flex justify-center" asChild>
                    <Button type="submit" className="w-1/3 m-auto">
                        <Text size={"3"}>{t("submit")}</Text>
                    </Button>
                </Form.Submit>
            </Form.Root>
        </Card>
    );
}
