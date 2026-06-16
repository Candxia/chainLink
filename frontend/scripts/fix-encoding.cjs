const fs = require('fs');
const path = require('path');
const Iconv = require('iconv-lite');

// If iconv-lite is not available, use a simpler approach
const dir = path.join(__dirname, '..', 'src', 'views');

function fixFile(filePath) {
  let content = fs.readFileSync(filePath);
  
  // Check if it starts with UTF-8 BOM
  if (content[0] === 0xEF && content[1] === 0xBB && content[2] === 0xBF) {
    content = content.slice(3);
  }
  
  // Try to detect if content has valid UTF-8 Chinese
  const utf8Str = content.toString('utf-8');
  
  // Check for the U+FFFD replacement character (broken encoding)
  if (utf8Str.includes('\uFFFD')) {
    console.log(`Fixing (FFFD): ${filePath}`);
    // Read as latin-1 to preserve bytes, then interpret as GBK, write as UTF-8
    const latin1 = content.toString('latin1');
    const gbkBuffer = Buffer.from(latin1, 'binary');
    const fixed = gbkBuffer.toString('utf-8');
    fs.writeFileSync(filePath, fixed, 'utf-8');
    return true;
  }
  
  // Check if Chinese text looks garbled (high bytes that look wrong in UTF-8)
  // A simple heuristic: count replacement chars
  const replCount = (utf8Str.match(/\uFFFD/g) || []).length;
  if (replCount > 3) {
    console.log(`Fixing (${replCount} replacements): ${filePath}`);
    const latin1 = content.toString('latin1');
    const gbkBuffer = Buffer.from(latin1, 'binary');
    const fixed = gbkBuffer.toString('utf-8');
    fs.writeFileSync(filePath, fixed, 'utf-8');
    return true;
  }
  
  return false;
}

function walkDir(dirPath) {
  const entries = fs.readdirSync(dirPath, { withFileTypes: true });
  for (const entry of entries) {
    const fullPath = path.join(dirPath, entry.name);
    if (entry.isDirectory()) {
      walkDir(fullPath);
    } else if (entry.name.endsWith('.vue') || entry.name.endsWith('.js')) {
      fixFile(fullPath);
    }
  }
}

walkDir(dir);
console.log('Done fixing files');
