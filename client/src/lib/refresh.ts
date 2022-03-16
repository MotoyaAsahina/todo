const t: Array<() => Promise<void>> = []
const tv: Array<() => void> = []

export const addRefreshListener = (...f: (() => Promise<void>)[]) => {
  t.push(...f)
}

export const addRefreshVoidListener = (...f: (() => void)[]) => {
  tv.push(...f)
}

export const refresh = async () => {
  for (const f of tv) {
    f()
  }
  for (const f of t) {
    await f()
  }
}
