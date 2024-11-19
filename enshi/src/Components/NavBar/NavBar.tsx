import {
    Container,
} from "@radix-ui/themes";
import SearchField from "./SearchField/SearchField";
import UserButton from "./UserButton/UserButton";
import CustomNavigationMenu from "./NavigationMenu/NavigationMenu";

export default function NavBar() {
    return (
        <Container size={"4"}>
            <nav className="flex justify-center pt-2 pb-2 ml-4 mr-4">
                <CustomNavigationMenu />

                <SearchField />

                <UserButton />
            </nav>
        </Container>
    );
}