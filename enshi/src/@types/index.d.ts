type TToast = {
    title: string;
    description?: string;
    action?: React.Component;
};

type TExistingToast = TToast & {
    id: number;
    resetFunc: (arg0: boolean) => void;
    open: boolean;
};
