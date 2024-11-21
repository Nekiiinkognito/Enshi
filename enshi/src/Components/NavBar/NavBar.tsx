import CustomNavigationMenu from "./NavigationMenu/NavigationMenu";
import RightButtonBar from "./RightButtonBar/RightButtonBar";
import SearchField from "./SearchField/SearchField";

export default function NavBar() {
    return (
        // <Container size={"4"}>
            <nav className="flex justify-center pt-2 pb-2 ml-4 mr-4">
                <CustomNavigationMenu />

                <SearchField />

                <RightButtonBar />
            </nav>
        // </Container>
    );
}