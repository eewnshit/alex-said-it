import Link from "next/link";
import Text from "./text";

export default function Navbar() {
    return (
        <div className="text-center p-3">
            <Link href={"/"}>
                <Text
                    className="cursor-pointer hover:opacity-20 duration-200"
                    display="inline"
                    style="italic"
                    weight="bold"
                    size={14}
                >
                    alex said it.
                </Text>
            </Link>
        </div>
    )
}
