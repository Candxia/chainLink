import mitt from 'mitt'
export const emitter = mitt()
export const useEventBus = () => emitter
