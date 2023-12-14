import { program } from 'commander';
import process from 'process';

program
    .version('0.0.1')
    .option('-t, --type <type>', 'Prompt type')
    .option('-i, --input-type', 'Input type')
    .option('-p, --prompt <prompt>', 'Prompt text')
    .parse(process.argv);

const options = program.opts();
const type = options.type;
const prompt = options.prompt;
const inputType = options.inputType;

if (type === undefined || prompt === undefined || inputType === undefined) {
    console.error("Missing required arguments");
    process.exit(1);
}

console.log(`Type: ${type}`);
console.log(`Input type: ${inputType}`)
console.log(`Prompt: ${prompt}`);