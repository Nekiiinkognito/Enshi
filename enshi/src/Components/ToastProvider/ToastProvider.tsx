import { Cross1Icon } from "@radix-ui/react-icons";
import * as Toast from "@radix-ui/react-toast";
import { Card, Text } from "@radix-ui/themes";
import { useAtomValue } from "jotai";
import React from "react";
import { toastAtom } from "../../AtomStore/AtomStore";

export default function ToastProvider(props: React.PropsWithChildren) {
    const toastsToRender = useAtomValue(toastAtom);

    return (
        <Toast.Provider swipeDirection="right">
            {props.children}

                {toastsToRender.map((toast) => {
                    return (
                        <Toast.Root
                            key={toast.id}
                            className="mt-2 mr-2 data-[state=open]:animate-slideFromRight data-[state=closed]:animate-fadeOut"
                            open={toast.open}
                            onOpenChange={toast.resetFunc}
                            color={"red"}
                        >
                            <Card className="relative w-60">
                                <Toast.Title>
                                    <Text size={"4"} weight={"bold"}>
                                        {toast.title}
                                    </Text>
                                </Toast.Title>
                                <Toast.Description className="overflow-hidden w-50 text-ellipsis">
                                    <Text className="text-pretty line-clamp-2">{toast.description}</Text>
                                </Toast.Description>
                                <Toast.Close className="absolute top-2 right-2">
                                    <Cross1Icon color="red" />
                                </Toast.Close>
                            </Card>
                        </Toast.Root>
                    );
                })}

            <Toast.Viewport className="fixed right-0 bottom-2" />
        </Toast.Provider>
    );
}
