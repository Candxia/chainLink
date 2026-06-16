import dayjs from 'dayjs'
export const formatToDateTzTime = (ts, fmt) => dayjs(ts).format(fmt || 'YYYY-MM-DD HH:mm:ss')
