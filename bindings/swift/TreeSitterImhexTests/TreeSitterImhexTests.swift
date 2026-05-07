import XCTest
import SwiftTreeSitter
import TreeSitterImhex

final class TreeSitterImhexTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_imhex())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Imhex grammar")
    }
}
