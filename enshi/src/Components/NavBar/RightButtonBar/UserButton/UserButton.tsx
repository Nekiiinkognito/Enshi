import {
    EnterIcon,
    ExitIcon,
    LaptopIcon,
    PersonIcon,
} from "@radix-ui/react-icons";
import { DropdownMenu, Flex, IconButton, Text } from "@radix-ui/themes";
import { Icon } from "@radix-ui/themes/dist/esm/components/callout.js";
import { useAtomValue } from "jotai";
import { useTranslation } from "react-i18next";
import { Link } from "react-router-dom";
import { userAtom } from "../../../../AtomStore/AtomStore";

export default function UserButton() {
    const user = useAtomValue(userAtom);
    const { t } = useTranslation();

    return (
        <div className="">
            <DropdownMenu.Root>
                <DropdownMenu.Trigger>
                    <IconButton>
                        <PersonIcon />
                    </IconButton>
                </DropdownMenu.Trigger>

                <DropdownMenu.Content className="w-fit">
                    <DropdownMenu.Item>
                        <Link to={"/user/:user-id/profile"}>
                            <Flex className="justify-between gap-2">
                                <Icon>
                                    <PersonIcon />
                                </Icon>

                                <Text>{t("profile")}</Text>
                            </Flex>
                        </Link>
                    </DropdownMenu.Item>

                    <DropdownMenu.Item>
                        <Link to={"/user/blogs"}>
                            <Flex className="justify-between gap-2">
                                <Icon>
                                    <LaptopIcon />
                                </Icon>
                                <Text>{t("yourBlogs")}</Text>
                            </Flex>
                        </Link>
                    </DropdownMenu.Item>

                    <DropdownMenu.Separator />

                    <DropdownMenu.Item color={user ? "red" : "green"}>
                        {user ? (
                            <Flex className="justify-between gap-2">
                                <Icon>
                                    <ExitIcon />
                                </Icon>
                                <Text>{t("signOut")}</Text>
                            </Flex>
                        ) : (
                            <Link to={"/login"}>
                                <Flex className="justify-between gap-2">
                                    <Icon>
                                        <EnterIcon />
                                    </Icon>
                                    <Text>{t("signIn")}</Text>
                                </Flex>
                            </Link>
                        )}
                    </DropdownMenu.Item>
                </DropdownMenu.Content>
            </DropdownMenu.Root>
        </div>
    );
}
