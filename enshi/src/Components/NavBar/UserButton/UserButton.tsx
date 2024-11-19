import {
    EnterIcon,
    ExitIcon,
    LaptopIcon,
    PersonIcon,
} from "@radix-ui/react-icons";
import { DropdownMenu, Flex, IconButton, Text } from "@radix-ui/themes";
import { Icon } from "@radix-ui/themes/dist/esm/components/callout.js";
import { useAtomValue } from "jotai";
import { userAtom } from "../../../AtomStore/AtomStore";
import { Link } from "react-router-dom";

export default function UserButton() {
    const user = useAtomValue(userAtom);

    return (
        <div className="flex justify-end flex-1">
            <DropdownMenu.Root>
                <DropdownMenu.Trigger>
                    <IconButton>
                        <PersonIcon />
                    </IconButton>
                </DropdownMenu.Trigger>

                <DropdownMenu.Content className="w-fit">
                    <DropdownMenu.Item>
                        <Link to={"/profile"}>
                            <Flex className="justify-between gap-2">
                                <Icon>
                                    <PersonIcon />
                                </Icon>

                                <Text>Profile</Text>
                            </Flex>
                        </Link>
                    </DropdownMenu.Item>

                    <DropdownMenu.Item>
                        <Flex className="justify-between gap-2">
                            <Icon>
                                <LaptopIcon />
                            </Icon>
                            <Text>Your blogs</Text>
                        </Flex>
                    </DropdownMenu.Item>

                    <DropdownMenu.Separator />

                    <DropdownMenu.Item color={user ? "red" : "green"}>
                        {user ? (
                            <Flex className="justify-between gap-2">
                                <Icon>
                                    <ExitIcon />
                                </Icon>
                                <Text>Log out</Text>
                            </Flex>
                        ) : (
                            <Flex className="justify-between gap-2">
                                <Icon>
                                    <EnterIcon />
                                </Icon>
                                <Text>Log in</Text>
                            </Flex>
                        )}
                    </DropdownMenu.Item>
                </DropdownMenu.Content>
            </DropdownMenu.Root>
        </div>
    );
}
