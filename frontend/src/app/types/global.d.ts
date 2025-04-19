declare module '*.woff' {
  const src: string

  export default src
}

declare module '*.woff2' {
  const src: string

  export default src
}

declare module '*.svg' {
  import * as React from 'react'
  export const ReactComponent: React.FunctionComponent<React.SVGProps<SVGSVGElement> & { title?: string }>
  const src: string
  export default src
}

declare module '*.svg?react' {
  export const ReactComponent: React.FunctionComponent<React.SVGProps<SVGSVGElement> & { title?: string }>
}

declare module '*.png'
declare module '*.gif'
declare module '*.jpg'
declare module '*.jpeg'
