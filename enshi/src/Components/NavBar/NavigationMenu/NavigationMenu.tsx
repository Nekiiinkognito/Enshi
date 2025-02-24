import * as NavigationMenu from "@radix-ui/react-navigation-menu";
import { Button, Heading, useThemeContext } from "@radix-ui/themes";
import { useTranslation } from "react-i18next";
import { useLocation, useNavigate } from "react-router-dom";

export default function CustomNavigationMenu() {

    const {t} = useTranslation()

    return (
        <div className="flex-1">
            <NavigationMenu.Root orientation="horizontal">
                <NavigationMenu.List className="flex items-center justify-start gap-8">
                    <NavItem text={t("home")} to="/" />

                    <NavItem text={t("following")} to="/c" />
                </NavigationMenu.List>
            </NavigationMenu.Root>
        </div>
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
