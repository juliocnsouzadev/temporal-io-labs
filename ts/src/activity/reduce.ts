import {Mapped} from "./map";
import {log} from "@temporalio/workflow";

export class Reduced {
    wordCounts: any;

    constructor() {
        this.wordCounts = {};
    }

    addWord(word: string) {
        let count = this.wordCounts[word];
        if (count === undefined) {
            count = 0;
        }
        this.wordCounts[word] = count + 1;
    }

}

export async function reduce(mapped:Mapped): Promise<Reduced> {
    const reduced:Reduced = new Reduced();
    for (const word of mapped.words) {
        reduced.addWord(word);
    }
    return reduced;
}