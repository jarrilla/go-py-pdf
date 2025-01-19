from flask import Flask, request, jsonify
import pymupdf4llm
import fitz
from io import BytesIO

app = Flask(__name__)

@app.route("/parse_pdf", methods=["POST"])
def extract_tables_endpoint():
    if not request.data:
        return jsonify({"error": "No PDF data provided"}), 400
    
    # Check Content-Type header
    content_type = request.headers.get('Content-Type', '')
    if content_type != 'application/pdf':
        return jsonify({"error": "Content-Type must be application/pdf"}), 415
    
    # Verify PDF signature
    if not request.data.startswith(b'%PDF-'):
        return jsonify({"error": "Invalid PDF format: File does not start with PDF signature"}), 415
    
    # Create BytesIO directly from request data
    pdf_bytes = BytesIO(request.data)
    
    result = {}
    try:
        # Create fitz Document from bytes
        doc = fitz.Document(stream=pdf_bytes)
        # Pass file to PyMuPDF
        md_text = pymupdf4llm.to_markdown(doc)
        result["markdown"] = md_text
    except Exception as e:
        return jsonify({"error": str(e)}), 500
    
    return jsonify(result)

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000, debug=True)