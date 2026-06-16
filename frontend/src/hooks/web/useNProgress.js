import nProgress from 'nprogress'
export const useNProgress = () => ({
  start: () => nProgress.start(),
  done: () => nProgress.done()
})
