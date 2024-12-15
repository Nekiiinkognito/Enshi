const isBigNumber = (num: any) => !Number.isSafeInteger(+num);

const enquoteBigNumber = (jsonString: any, bigNumChecker: any) =>
    jsonString.replaceAll(
        /([:\s\[,]*)(\d+)([\s,\]]*)/g,
        (matchingSubstr: any, prefix: any, bigNum: any, suffix: any) =>
            bigNumChecker(bigNum)
                ? `${prefix}"${bigNum}"${suffix}`
                : matchingSubstr
    );

const parseWithBigInt = (jsonString: any, bigNumChecker: any) =>
    JSON.parse(enquoteBigNumber(jsonString, bigNumChecker), (_key, value) =>
        !isNaN(value) && bigNumChecker(value) ? BigInt(value).toString() : value
    );

export const JSONWithInt64 = (jsonString: any) => {
    return parseWithBigInt(jsonString, isBigNumber);
};
