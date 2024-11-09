import Quill, { Delta,  } from "quill/core";
import ReactQuill from "react-quill";
import React, {
    forwardRef,
    useEffect,
    useLayoutEffect,
    useRef,
    useState,
} from "react";
import Sources from "quill";

type TEditor = {
    readOnly?: boolean;
    defaultValue?: string | Delta;
    onChange: (d: string) => void; // TODO: make type
    onSelectionChange?: any; // TODO same as before
};

const modules = {
    toolbar: [
        [{ header: [1, 2, 3, false] }],
        ["bold", "italic", "underline", "strike", "blockquote"],
        [
            { list: "ordered" },
            { list: "bullet" },
            { indent: "-1" },
            { indent: "+1" },
        ],
        ["link", "image"],
        ["clean"],
        [{ align: [] }],
    ],
};

/**
 * @param onChange - function that accepts Delta element
 */
const Editor = forwardRef((props: TEditor) => {
    const editor = useRef(null);
    const [quill, setQuill] = useState<Quill | null>(null);
    const [value, setValue] = useState(new Delta())

    useEffect(() => {
        if (editor.current) {
            //@ts-ignore
            const temp = editor.current.getEditor() as Quill;
            setQuill(temp);
        }
        return () => {
            setQuill(null);
        };
    }, [editor.current]);

    const changeHandler = (val: string, _changeDelta: Delta, _source: Sources, _editor: ReactQuill.UnprivilegedEditor) => {
        console.log(val);
        console.log(JSON.stringify(quill?.getContents().ops, null, 2))
        let fullDelta = quill?.getContents()
        props.onChange(val || "")
        setValue(fullDelta || new Delta())
    }

    return (
        <div className="text-editor">
            <ReactQuill
                value={value}
                ref={editor}
                modules={modules}


                onChange={changeHandler}
                
                
                theme="snow"
                placeholder="Type your thoughts here..."
            />
        </div>
    );
});

export default Editor;
