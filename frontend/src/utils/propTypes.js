export const propTypes = new Proxy({}, { get: () => ({ def: () => undefined }) })
