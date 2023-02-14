import algosdk from "algosdk";

try {
  let count = 0;
  let t0 = Date.now();
  const list : string[] = [];
  while (count < 5) {
    const { addr } = algosdk.generateAccount();
    if (addr.startsWith("TEST")) list[count++] = addr;
  }
  console.log(`JS Found 5 adresses in ${Math.floor((Date.now() - t0) / 1000)} secs\n${list}\n`);
} catch (error) {
  console.log((error as Error).message);
}
