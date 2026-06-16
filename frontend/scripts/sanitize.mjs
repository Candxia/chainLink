import { readFileSync, writeFileSync } from 'fs'
import { join, dirname } from 'path'
import { fileURLToPath } from 'url'

const __dirname = dirname(fileURLToPath(import.meta.url))
const base = join(__dirname, '../src')

// Find all Vue/TS/TSX files and fix encoding issues
// Problem: Some Chinese chars in source files became corrupted during copy
// Fix: Strip all non-ASCII chars from .vue and .tsx files (leave only valid UTF-8)

function fixStringLiterals(fp) {
  let raw = readFileSync(fp, 'utf-8')
  const orig = raw
  
  // Strip control chars and replacement chars
  raw = raw.replace(/\uFFFD/g, '')
  
  // Fix common corruption: `label='xxx\?' or `label="xxx\?"` where `\?` is U+FFFD followed by ?
  raw = raw.replace(/(label=["'])[^"']*\uFFFD*[?"']*/g, '$1')  
  raw = raw.replace(/(placeholder=["'])[^"']*\uFFFD*[?"']*/g, '$1')
  
  // Remove lines that contain only `// ` with corrupted Chinese (empty comments)
  const lines = raw.split('\n')
  const cleaned = lines.map(line => {
    // If the line has `// ` followed by non-ASCII garbage, clear the comment
    if (/\/\/ [\u0080-\ufffd]/.test(line) && !/[a-zA-Z0-9]/.test(line.replace(/\/\/ /, ''))) {
      return line.replace(/\/\/.*/, '')
    }
    return line
  })
  raw = cleaned.join('\n')
  
  if (raw !== orig) {
    writeFileSync(fp, raw, 'utf-8')
    return true
  }
  return false
}

function walk(dir) {
  const { readdirSync } = await import('fs')
  const entries = readdirSync(dir, { withFileTypes: true })
  for (const e of entries) {
    const p = join(dir, e.name)
    if (e.isDirectory()) walk(p)
    else if (/\.(vue|tsx?|jsx?)$/.test(e.name)) fixStringLiterals(p)
  }
}

walk(base)
console.log('All files sanitized')
