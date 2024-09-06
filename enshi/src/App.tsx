import "./App.css";
import "@radix-ui/themes/styles.css";
import {
    Badge,
    Box,
    Button,
    Callout,
    Flex,
    Text,
    Theme,
    ThemePanel,
} from "@radix-ui/themes";

import React, { PropsWithChildren } from "react";
import * as Tooltip from "@radix-ui/react-tooltip";

const MyButton = React.forwardRef<HTMLButtonElement, PropsWithChildren>((props, forwardedRef) => (
  <Button {...props} ref={forwardedRef} />
));

function App() {
    return (
        <Theme className="h-fit" accentColor="indigo" grayColor="slate">
            <Flex className="w-full">
                <Callout.Root>
                    <Text>Hello world!</Text>
                    <Button highContrast variant="soft">
                        Let's goooo!
                    </Button>
                    <Button variant="soft">Let's goooo!</Button>
                    <Badge color="red">What</Badge>
                </Callout.Root>
            </Flex>

            
                <Tooltip.Root>
                    <Tooltip.Trigger>
                            <MyButton>Open dialog</MyButton>
                    </Tooltip.Trigger>
                    <Tooltip.Portal>
                      <Tooltip.Content>
                        apsofjiosfwiuhfwei
                      </Tooltip.Content>
                    </Tooltip.Portal>
                </Tooltip.Root>

            <ThemePanel />
        </Theme>
    );
}

export default App;
