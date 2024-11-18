import { Button, Heading, IconButton, TextField, useThemeContext } from "@radix-ui/themes";
import * as NavigationMenu from "@radix-ui/react-navigation-menu";
import { useLocation, useNavigate } from "react-router-dom";
import { PersonIcon } from "@radix-ui/react-icons"
 
export default function NavBar() {
    return (
        <nav className="flex justify-between pt-2 pb-2 ml-4 mr-4">
            <NavigationMenu.Root orientation="horizontal">
                <NavigationMenu.List className="flex items-center justify-start gap-8">
                    <NavItem text="Home" to="/" />

                    <NavItem text="Following" to="/c" />
                </NavigationMenu.List>
            </NavigationMenu.Root>

            <TextField.Root className="w-1/3 rounded-lg" placeholder="Search...">
            </TextField.Root>

            <IconButton>
                <PersonIcon />
            </IconButton>
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
    const theme = useThemeContext();

    return (
        <div className="relative flex flex-col">
            <NavigationMenu.Item>
                <NavigationMenu.Link>
                    <div>
                        <Button
                            className="w-fit border-0 border-b-[0px] border-solid"
                            highContrast
                            variant="ghost"
                            onClick={() => navigate(props.to)}
                        >
                            <Heading weight={"medium"} size={"3"}>
                                {props.text}
                            </Heading>
                        </Button>
                    </div>
                </NavigationMenu.Link>
            </NavigationMenu.Item>
            {location.pathname == props.to ? (
                <div
                    className={`absolute animate-widthOut bottom-[-0.35rem] 
                                w-full h-[2px] z-[999] rounded-full`}
                    style={{
                        background: `var(--${theme.accentColor}-10)`,
                    }}
                ></div>
            ) : null}
        </div>
    );
}
