import { Button, Text } from "@radix-ui/themes";
import * as NavigationMenu from "@radix-ui/react-navigation-menu";
import { Link, useLocation, useNavigate } from "react-router-dom";

export default function NavBar() {
    return (
        <nav className="pt-2">
            <NavigationMenu.Root
                orientation="horizontal"
                className="flex justify-center"
            >
                <NavigationMenu.List className="flex justify-center gap-2">
                    <NavItem text="Cringer" to="/"/>

                    <NavItem text="C-Cringer" to="/c"/>

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

    const navigate = useNavigate()
    const location = useLocation()

    console.log(location);
    

    return (
        <NavigationMenu.Item>
            <NavigationMenu.Link>
                    <Button
                        className="w-fit h-fit rounded-full m-0 p-0 pr-2 pl-2 mt-2 mb-2"

                        highContrast

                        variant={location.pathname === props.to ? 'solid' : 'ghost'}
                        onClick={() => navigate(props.to)}
                    >
                        <Text size={"3"}>{props.text}</Text>
                    </Button>
            </NavigationMenu.Link>
        </NavigationMenu.Item>
    );
}
