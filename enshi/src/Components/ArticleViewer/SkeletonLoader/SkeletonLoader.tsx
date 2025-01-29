import { Container, Skeleton, Text } from "@radix-ui/themes";
import {
    headerLong,
    headerShort,
    pText,
} from "../../../constants/textForSkeleton";

export default function SkeletonPostLoader() {
    return (
        <Container size={"2"} className="mt-4">
            <Skeleton>
                <Text size={"6"}>{headerLong}</Text>
                <br />
                <Text size={"6"}>{headerShort}</Text>
                <br />
                <br />
                <Text>{pText}</Text>
                <br />
                <br />
                <Text wrap={"pretty"}>{pText}</Text>
                <br />
                <br />
                <Text>{pText}</Text>
            </Skeleton>
        </Container>
    );
}
