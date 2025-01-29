import { Box, Skeleton } from "@radix-ui/themes";

export default function SkeletonBoxes() {
    return (
        <>
            <Skeleton>
                <Box className="w-full h-20 mb-2 rounded-lg"></Box>
            </Skeleton>
            <Skeleton>
                <Box className="w-full h-20 mb-2 rounded-lg"></Box>
            </Skeleton>
            <Skeleton>
                <Box className="w-full h-20 mb-2 rounded-lg"></Box>
            </Skeleton>
        </>
    );
}
