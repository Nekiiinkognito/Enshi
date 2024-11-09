import { Button, Card, ChevronDownIcon, Text } from "@radix-ui/themes";
import * as NavigationMenu from "@radix-ui/react-navigation-menu";
import { useLocation, useNavigate } from "react-router-dom";

export default function NavBar() {
    return (
        <nav className="pt-2">
            <NavigationMenu.Root
                orientation="horizontal"
                className="flex justify-center"
            >
                <NavigationMenu.List className="flex justify-center gap-2">
                    <NavItem text="Cringer" to="/" />

                    <NavItem text="C-Cringer" to="/c" />

                    <NavigationMenu.Item className="text-center">
                        <NavigationMenu.Trigger className="flex items-center">
                            <Button
                                asChild
                                className="w-fit pr-2 h-fit rounded-full m-0 p-0  pl-2 mt-2 mb-2 duration-[50ms]"
                                variant="ghost"
                                highContrast
                            >
                                <Text
                                    size={"3"}
                                    className="flex items-center gap-1"
                                >
                                    Cringer 123 <ChevronDownIcon />
                                </Text>
                            </Button>
                        </NavigationMenu.Trigger>

                        <NavigationMenu.Content className="absolute data-[motion=from-start]:scale-150">
                            <Card>asd</Card>
                        </NavigationMenu.Content>
                    </NavigationMenu.Item>
                </NavigationMenu.List>
            </NavigationMenu.Root>
        </nav>
    );
}

type TNavItem = {
    text: string;
    to: string;
};

function NavItem(props: TNavItem) {
    const navigate = useNavigate();
    const location = useLocation();

    return (
        <NavigationMenu.Item>
            <NavigationMenu.Link>
                <Button
                    className="w-fit h-fit rounded-full m-0 p-0 pr-2 pl-2 mt-2 mb-2 duration-[50ms]"
                    highContrast
                    variant={location.pathname === props.to ? "solid" : "ghost"}
                    onClick={() => navigate(props.to)}
                >
                    <Text size={"3"}>{props.text}</Text>
                </Button>
            </NavigationMenu.Link>
        </NavigationMenu.Item>
    );
}
